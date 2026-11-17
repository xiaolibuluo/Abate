
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK478_0(v) (((v) << 4) | ((v) & 0x0F))

int classify478_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK478_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd478_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale478_1(int value) {
    return gcd478_1(value, 11) * 18;
}
#define MASK478_2(v) (((v) << 4) | ((v) & 0x0F))

int classify478_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK478_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK478_3(v) (((v) << 4) | ((v) & 0x0F))

int classify478_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK478_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind478_4 { KIND_EMPTY478_4 = 0, KIND_FULL478_4 = 39 };

union Cell478_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind478_4(union Cell478_4 *cell) {
    cell->number = 44;
    return cell->number > 0 ? KIND_FULL478_4 : KIND_EMPTY478_4;
}
enum Kind478_5 { KIND_EMPTY478_5 = 0, KIND_FULL478_5 = 46 };

union Cell478_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind478_5(union Cell478_5 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL478_5 : KIND_EMPTY478_5;
}
struct Node478_6 {
    int value;
    struct Node478_6 *next;
};

int length478_6(const struct Node478_6 *head) {
    int count = 0;
    const struct Node478_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 53;
}
