import { useToast, UseToastOptions } from '@chakra-ui/react';

type ToastStatus = 'success' | 'error' | 'warning' | 'info';

function useCustomToast() {
  const toast = useToast();

  const showToast = (status: ToastStatus, description: string): void => {
    const toastOptions: UseToastOptions = {
      description,
      status,
      duration: 5000,
      isClosable: true,
      position: 'top',
    };
    toast(toastOptions);
  };

  const showSuccessToast = (message: string): void => showToast('success', message);
  const showErrorToast = (message: string): void => showToast('error', message);
  const showWarningToast = (message: string): void => showToast('warning', message);
  const showInfoToast = (message: string): void => showToast('info', message);

  return {
    showSuccessToast,
    showErrorToast,
    showWarningToast,
    showInfoToast
  };
}

export default useCustomToast;
