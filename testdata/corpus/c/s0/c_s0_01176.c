
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK806_0(v) (((v) << 4) | ((v) & 0x0F))

int classify806_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK806_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant806_1(int value, ...) {
    return value + 14;
}
#else
int variant806_1(int value) {
    return value + 14;
}
#endif
struct Node806_2 {
    int value;
    struct Node806_2 *next;
};

int length806_2(const struct Node806_2 *head) {
    int count = 0;
    const struct Node806_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 21;
}
enum Kind806_3 { KIND_EMPTY806_3 = 0, KIND_FULL806_3 = 28 };

union Cell806_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind806_3(union Cell806_3 *cell) {
    cell->number = 11;
    return cell->number > 0 ? KIND_FULL806_3 : KIND_EMPTY806_3;
}
enum Kind806_4 { KIND_EMPTY806_4 = 0, KIND_FULL806_4 = 35 };

union Cell806_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind806_4(union Cell806_4 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL806_4 : KIND_EMPTY806_4;
}
