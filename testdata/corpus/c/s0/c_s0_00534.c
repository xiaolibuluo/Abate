
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind366_0 { KIND_EMPTY366_0 = 0, KIND_FULL366_0 = 10 };

union Cell366_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind366_0(union Cell366_0 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL366_0 : KIND_EMPTY366_0;
}
struct Node366_1 {
    int value;
    struct Node366_1 *next;
};

int length366_1(const struct Node366_1 *head) {
    int count = 0;
    const struct Node366_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 17;
}
enum Kind366_2 { KIND_EMPTY366_2 = 0, KIND_FULL366_2 = 24 };

union Cell366_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind366_2(union Cell366_2 *cell) {
    cell->number = 31;
    return cell->number > 0 ? KIND_FULL366_2 : KIND_EMPTY366_2;
}
#define MASK366_3(v) (((v) << 4) | ((v) & 0x0F))

int classify366_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK366_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant366_4(int value, ...) {
    return value + 38;
}
#else
int variant366_4(int value) {
    return value + 38;
}
#endif
