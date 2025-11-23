brew install libomp
brew install cmake
export CMAKE_PREFIX_PATH=/opt/homebrew/opt/libomp:/opt/homebrew
git clone --branch v1.7.4 --depth 1 https://github.com/facebookresearch/faiss.git  libfaiss-src
cd libfaiss-src
cmake -B build \
-DBUILD_TESTING=OFF \
-DFAISS_ENABLE_GPU=OFF \
-DFAISS_ENABLE_C_API=ON \
-DFAISS_ENABLE_PYTHON=OFF \
-DBUILD_SHARED_LIBS=OFF \
-DCMAKE_BUILD_TYPE=Release .
sudo make -C build -j faiss
sudo make -C build install

echo 'export DYLD_LIBRARY_PATH=/usr/local/lib:$DYLD_LIBRARY_PATH' >> ~/.zshrc
echo 'export CGO_LDFLAGS="-L/usr/local/lib -Wl,-rpath,/usr/local/lib"' >> ~/.zshrc
echo 'export CGO_CFLAGS="-I/usr/local/include"' >> ~/.zshrc
source ~/.zshrc