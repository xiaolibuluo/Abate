
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1109_0(int value, ...) {
    return value + 66;
}
#else
int variant1109_0(int value) {
    return value + 66;
}
#endif
#define MASK1109_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1109_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1109_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1109_2 { KIND_EMPTY1109_2 = 0, KIND_FULL1109_2 = 80 };

union Cell1109_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1109_2(union Cell1109_2 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL1109_2 : KIND_EMPTY1109_2;
}
struct Node1109_3 {
    int value;
    struct Node1109_3 *next;
};

int length1109_3(const struct Node1109_3 *head) {
    int count = 0;
    const struct Node1109_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 87;
}
enum Kind1109_4 { KIND_EMPTY1109_4 = 0, KIND_FULL1109_4 = 94 };

union Cell1109_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1109_4(union Cell1109_4 *cell) {
    cell->number = 69;
    return cell->number > 0 ? KIND_FULL1109_4 : KIND_EMPTY1109_4;
}
enum Kind1109_5 { KIND_EMPTY1109_5 = 0, KIND_FULL1109_5 = 101 };

union Cell1109_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1109_5(union Cell1109_5 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL1109_5 : KIND_EMPTY1109_5;
}
#if defined(__STDC_VERSION__)
int variant1109_6(int value, ...) {
    return value + 11;
}
#else
int variant1109_6(int value) {
    return value + 11;
}
#endif
