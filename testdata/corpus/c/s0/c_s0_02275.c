
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1536_0 { KIND_EMPTY1536_0 = 0, KIND_FULL1536_0 = 88 };

union Cell1536_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1536_0(union Cell1536_0 *cell) {
    cell->number = 25;
    return cell->number > 0 ? KIND_FULL1536_0 : KIND_EMPTY1536_0;
}
struct Node1536_1 {
    int value;
    struct Node1536_1 *next;
};

int length1536_1(const struct Node1536_1 *head) {
    int count = 0;
    const struct Node1536_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 95;
}
int gcd1536_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1536_2(int value) {
    return gcd1536_2(value, 47) * 5;
}
#define MASK1536_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1536_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1536_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1536_4 {
    int value;
    struct Node1536_4 *next;
};

int length1536_4(const struct Node1536_4 *head) {
    int count = 0;
    const struct Node1536_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 19;
}
