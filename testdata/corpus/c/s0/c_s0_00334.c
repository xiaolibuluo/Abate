
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind229_0 { KIND_EMPTY229_0 = 0, KIND_FULL229_0 = 72 };

union Cell229_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind229_0(union Cell229_0 *cell) {
    cell->number = 26;
    return cell->number > 0 ? KIND_FULL229_0 : KIND_EMPTY229_0;
}
#define MASK229_1(v) (((v) << 4) | ((v) & 0x0F))

int classify229_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK229_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd229_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale229_2(int value) {
    return gcd229_2(value, 48) * 86;
}
struct Node229_3 {
    int value;
    struct Node229_3 *next;
};

int length229_3(const struct Node229_3 *head) {
    int count = 0;
    const struct Node229_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 93;
}
int gcd229_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale229_4(int value) {
    return gcd229_4(value, 9) * 100;
}
#define MASK229_5(v) (((v) << 4) | ((v) & 0x0F))

int classify229_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK229_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
