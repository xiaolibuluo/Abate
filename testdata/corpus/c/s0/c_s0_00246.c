
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant170_0(int value, ...) {
    return value + 81;
}
#else
int variant170_0(int value) {
    return value + 81;
}
#endif
enum Kind170_1 { KIND_EMPTY170_1 = 0, KIND_FULL170_1 = 88 };

union Cell170_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind170_1(union Cell170_1 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL170_1 : KIND_EMPTY170_1;
}
#define MASK170_2(v) (((v) << 4) | ((v) & 0x0F))

int classify170_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK170_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node170_3 {
    int value;
    struct Node170_3 *next;
};

int length170_3(const struct Node170_3 *head) {
    int count = 0;
    const struct Node170_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 5;
}
int gcd170_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale170_4(int value) {
    return gcd170_4(value, 23) * 12;
}
