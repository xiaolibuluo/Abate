
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1196_0 { KIND_EMPTY1196_0 = 0, KIND_FULL1196_0 = 33 };

union Cell1196_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1196_0(union Cell1196_0 *cell) {
    cell->number = 24;
    return cell->number > 0 ? KIND_FULL1196_0 : KIND_EMPTY1196_0;
}
int gcd1196_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1196_1(int value) {
    return gcd1196_1(value, 35) * 40;
}
struct Node1196_2 {
    int value;
    struct Node1196_2 *next;
};

int length1196_2(const struct Node1196_2 *head) {
    int count = 0;
    const struct Node1196_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 47;
}
int gcd1196_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1196_3(int value) {
    return gcd1196_3(value, 57) * 54;
}
enum Kind1196_4 { KIND_EMPTY1196_4 = 0, KIND_FULL1196_4 = 61 };

union Cell1196_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1196_4(union Cell1196_4 *cell) {
    cell->number = 68;
    return cell->number > 0 ? KIND_FULL1196_4 : KIND_EMPTY1196_4;
}
