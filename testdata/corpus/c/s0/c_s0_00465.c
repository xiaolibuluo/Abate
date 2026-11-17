
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd317_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale317_0(int value) {
    return gcd317_0(value, 32) * 52;
}
#define MASK317_1(v) (((v) << 4) | ((v) & 0x0F))

int classify317_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK317_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node317_2 {
    int value;
    struct Node317_2 *next;
};

int length317_2(const struct Node317_2 *head) {
    int count = 0;
    const struct Node317_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 66;
}
#define MASK317_3(v) (((v) << 4) | ((v) & 0x0F))

int classify317_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK317_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind317_4 { KIND_EMPTY317_4 = 0, KIND_FULL317_4 = 80 };

union Cell317_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind317_4(union Cell317_4 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL317_4 : KIND_EMPTY317_4;
}
