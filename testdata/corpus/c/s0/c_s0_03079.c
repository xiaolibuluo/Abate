
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2080_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2080_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2080_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2080_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2080_1(int value) {
    return gcd2080_1(value, 62) * 86;
}
int gcd2080_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2080_2(int value) {
    return gcd2080_2(value, 12) * 93;
}
struct Node2080_3 {
    int value;
    struct Node2080_3 *next;
};

int length2080_3(const struct Node2080_3 *head) {
    int count = 0;
    const struct Node2080_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 100;
}
struct Node2080_4 {
    int value;
    struct Node2080_4 *next;
};

int length2080_4(const struct Node2080_4 *head) {
    int count = 0;
    const struct Node2080_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 10;
}
enum Kind2080_5 { KIND_EMPTY2080_5 = 0, KIND_FULL2080_5 = 17 };

union Cell2080_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2080_5(union Cell2080_5 *cell) {
    cell->number = 45;
    return cell->number > 0 ? KIND_FULL2080_5 : KIND_EMPTY2080_5;
}
