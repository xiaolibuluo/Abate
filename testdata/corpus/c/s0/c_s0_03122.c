
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2114_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2114_0(int value) {
    return gcd2114_0(value, 45) * 36;
}
struct Node2114_1 {
    int value;
    struct Node2114_1 *next;
};

int length2114_1(const struct Node2114_1 *head) {
    int count = 0;
    const struct Node2114_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 43;
}
enum Kind2114_2 { KIND_EMPTY2114_2 = 0, KIND_FULL2114_2 = 50 };

union Cell2114_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2114_2(union Cell2114_2 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL2114_2 : KIND_EMPTY2114_2;
}
#define MASK2114_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2114_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2114_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
