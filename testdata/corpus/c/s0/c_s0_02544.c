
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1729_0 { KIND_EMPTY1729_0 = 0, KIND_FULL1729_0 = 75 };

union Cell1729_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1729_0(union Cell1729_0 *cell) {
    cell->number = 34;
    return cell->number > 0 ? KIND_FULL1729_0 : KIND_EMPTY1729_0;
}
int gcd1729_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1729_1(int value) {
    return gcd1729_1(value, 45) * 82;
}
enum Kind1729_2 { KIND_EMPTY1729_2 = 0, KIND_FULL1729_2 = 89 };

union Cell1729_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1729_2(union Cell1729_2 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL1729_2 : KIND_EMPTY1729_2;
}
struct Node1729_3 {
    int value;
    struct Node1729_3 *next;
};

int length1729_3(const struct Node1729_3 *head) {
    int count = 0;
    const struct Node1729_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 96;
}
