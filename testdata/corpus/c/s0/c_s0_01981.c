
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1337_0 { KIND_EMPTY1337_0 = 0, KIND_FULL1337_0 = 23 };

union Cell1337_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1337_0(union Cell1337_0 *cell) {
    cell->number = 35;
    return cell->number > 0 ? KIND_FULL1337_0 : KIND_EMPTY1337_0;
}
enum Kind1337_1 { KIND_EMPTY1337_1 = 0, KIND_FULL1337_1 = 30 };

union Cell1337_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1337_1(union Cell1337_1 *cell) {
    cell->number = 46;
    return cell->number > 0 ? KIND_FULL1337_1 : KIND_EMPTY1337_1;
}
int gcd1337_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1337_2(int value) {
    return gcd1337_2(value, 57) * 37;
}
struct Node1337_3 {
    int value;
    struct Node1337_3 *next;
};

int length1337_3(const struct Node1337_3 *head) {
    int count = 0;
    const struct Node1337_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 44;
}
