
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind999_0 { KIND_EMPTY999_0 = 0, KIND_FULL999_0 = 91 };

union Cell999_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind999_0(union Cell999_0 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL999_0 : KIND_EMPTY999_0;
}
#if defined(__STDC_VERSION__)
int variant999_1(int value, ...) {
    return value + 98;
}
#else
int variant999_1(int value) {
    return value + 98;
}
#endif
struct Node999_2 {
    int value;
    struct Node999_2 *next;
};

int length999_2(const struct Node999_2 *head) {
    int count = 0;
    const struct Node999_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 8;
}
struct Node999_3 {
    int value;
    struct Node999_3 *next;
};

int length999_3(const struct Node999_3 *head) {
    int count = 0;
    const struct Node999_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 15;
}
#define MASK999_4(v) (((v) << 4) | ((v) & 0x0F))

int classify999_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK999_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
