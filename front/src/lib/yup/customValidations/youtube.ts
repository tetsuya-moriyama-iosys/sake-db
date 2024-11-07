import * as yup from 'yup';

const youtubeURLRegex =
  /^(https?:\/\/)?(www\.)?(youtube\.com|youtu\.be)\/(watch\?v=|embed\/|v\/|.+\?v=)?([^&=%?]{11})$/;

export const youtube = () => {
  return yup
    .string()
    .test(
      'youtube',
      'URLはYouTubeの動画リンク形式である必要があります',
      (value) => {
        if (!value) return true; // 値がない場合はスキップ（nullable対応）
        return youtubeURLRegex.test(value); // 正規表現でYouTube URLをチェック
      },
    );
};
