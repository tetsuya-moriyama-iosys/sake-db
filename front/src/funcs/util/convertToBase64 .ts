//FileReaderが非同期なのでPromiseで実装している
const convertToBase64 = (file: File): Promise<string> => {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onloadend = () => {
      resolve(reader.result?.toString().split(',')[1] || '');
    };
    reader.onerror = () => {
      reject(new Error('Error converting file to Base64'));
    };
    reader.readAsDataURL(file);
  });
};

export default convertToBase64;
