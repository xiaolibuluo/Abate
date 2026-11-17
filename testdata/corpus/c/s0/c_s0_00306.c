
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind211_0 { KIND_EMPTY211_0 = 0, KIND_FULL211_0 = 32 };

union Cell211_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind211_0(union Cell211_0 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL211_0 : KIND_EMPTY211_0;
}
#if defined(__STDC_VERSION__)
int variant211_1(int value, ...) {
    return value + 39;
}
#else
int variant211_1(int value) {
    return value + 39;
}
#endif
enum Kind211_2 { KIND_EMPTY211_2 = 0, KIND_FULL211_2 = 46 };

union Cell211_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind211_2(union Cell211_2 *cell) {
    cell->number = 44;
    return cell->number > 0 ? KIND_FULL211_2 : KIND_EMPTY211_2;
}
enum Kind211_3 { KIND_EMPTY211_3 = 0, KIND_FULL211_3 = 53 };

union Cell211_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind211_3(union Cell211_3 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL211_3 : KIND_EMPTY211_3;
}
