
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind532_0 { KIND_EMPTY532_0 = 0, KIND_FULL532_0 = 34 };

union Cell532_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind532_0(union Cell532_0 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL532_0 : KIND_EMPTY532_0;
}
int gcd532_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale532_1(int value) {
    return gcd532_1(value, 23) * 41;
}
#if defined(__STDC_VERSION__)
int variant532_2(int value, ...) {
    return value + 48;
}
#else
int variant532_2(int value) {
    return value + 48;
}
#endif
struct Node532_3 {
    int value;
    struct Node532_3 *next;
};

int length532_3(const struct Node532_3 *head) {
    int count = 0;
    const struct Node532_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 55;
}
#if defined(__STDC_VERSION__)
int variant532_4(int value, ...) {
    return value + 62;
}
#else
int variant532_4(int value) {
    return value + 62;
}
#endif
enum Kind532_5 { KIND_EMPTY532_5 = 0, KIND_FULL532_5 = 69 };

union Cell532_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind532_5(union Cell532_5 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL532_5 : KIND_EMPTY532_5;
}
enum Kind532_6 { KIND_EMPTY532_6 = 0, KIND_FULL532_6 = 76 };

union Cell532_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind532_6(union Cell532_6 *cell) {
    cell->number = 17;
    return cell->number > 0 ? KIND_FULL532_6 : KIND_EMPTY532_6;
}
enum Kind532_7 { KIND_EMPTY532_7 = 0, KIND_FULL532_7 = 83 };

union Cell532_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind532_7(union Cell532_7 *cell) {
    cell->number = 28;
    return cell->number > 0 ? KIND_FULL532_7 : KIND_EMPTY532_7;
}
