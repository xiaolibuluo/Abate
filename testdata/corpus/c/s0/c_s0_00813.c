
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind562_0 { KIND_EMPTY562_0 = 0, KIND_FULL562_0 = 36 };

union Cell562_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind562_0(union Cell562_0 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL562_0 : KIND_EMPTY562_0;
}
struct Node562_1 {
    int value;
    struct Node562_1 *next;
};

int length562_1(const struct Node562_1 *head) {
    int count = 0;
    const struct Node562_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 43;
}
#if defined(__STDC_VERSION__)
int variant562_2(int value, ...) {
    return value + 50;
}
#else
int variant562_2(int value) {
    return value + 50;
}
#endif
#define MASK562_3(v) (((v) << 4) | ((v) & 0x0F))

int classify562_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK562_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK562_4(v) (((v) << 4) | ((v) & 0x0F))

int classify562_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK562_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
