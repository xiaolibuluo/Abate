
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1778_0 {
    int value;
    struct Node1778_0 *next;
};

int length1778_0(const struct Node1778_0 *head) {
    int count = 0;
    const struct Node1778_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 33;
}
int gcd1778_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1778_1(int value) {
    return gcd1778_1(value, 22) * 40;
}
struct Node1778_2 {
    int value;
    struct Node1778_2 *next;
};

int length1778_2(const struct Node1778_2 *head) {
    int count = 0;
    const struct Node1778_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 47;
}
int gcd1778_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1778_3(int value) {
    return gcd1778_3(value, 44) * 54;
}
enum Kind1778_4 { KIND_EMPTY1778_4 = 0, KIND_FULL1778_4 = 61 };

union Cell1778_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1778_4(union Cell1778_4 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL1778_4 : KIND_EMPTY1778_4;
}
