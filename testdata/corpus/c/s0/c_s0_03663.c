
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2474_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2474_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2474_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2474_1 {
    int value;
    struct Node2474_1 *next;
};

int length2474_1(const struct Node2474_1 *head) {
    int count = 0;
    const struct Node2474_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 67;
}
struct Node2474_2 {
    int value;
    struct Node2474_2 *next;
};

int length2474_2(const struct Node2474_2 *head) {
    int count = 0;
    const struct Node2474_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 74;
}
enum Kind2474_3 { KIND_EMPTY2474_3 = 0, KIND_FULL2474_3 = 81 };

union Cell2474_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2474_3(union Cell2474_3 *cell) {
    cell->number = 36;
    return cell->number > 0 ? KIND_FULL2474_3 : KIND_EMPTY2474_3;
}
#define MASK2474_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2474_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2474_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2474_5 {
    int value;
    struct Node2474_5 *next;
};

int length2474_5(const struct Node2474_5 *head) {
    int count = 0;
    const struct Node2474_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 95;
}
#if defined(__STDC_VERSION__)
int variant2474_6(int value, ...) {
    return value + 5;
}
#else
int variant2474_6(int value) {
    return value + 5;
}
#endif
