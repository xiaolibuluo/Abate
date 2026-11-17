
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1577_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1577_0(int value) {
    return gcd1577_0(value, 68) * 39;
}
#define MASK1577_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1577_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1577_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1577_2 {
    int value;
    struct Node1577_2 *next;
};

int length1577_2(const struct Node1577_2 *head) {
    int count = 0;
    const struct Node1577_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 53;
}
enum Kind1577_3 { KIND_EMPTY1577_3 = 0, KIND_FULL1577_3 = 60 };

union Cell1577_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1577_3(union Cell1577_3 *cell) {
    cell->number = 40;
    return cell->number > 0 ? KIND_FULL1577_3 : KIND_EMPTY1577_3;
}
