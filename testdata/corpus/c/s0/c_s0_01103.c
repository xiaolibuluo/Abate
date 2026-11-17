
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan755_0(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 23;
}
struct Node755_1 {
    int value;
    struct Node755_1 *next;
};

int length755_1(const struct Node755_1 *head) {
    int count = 0;
    const struct Node755_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 30;
}
enum Kind755_2 { KIND_EMPTY755_2 = 0, KIND_FULL755_2 = 37 };

union Cell755_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind755_2(union Cell755_2 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL755_2 : KIND_EMPTY755_2;
}
int gcd755_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale755_3(int value) {
    return gcd755_3(value, 20) * 44;
}
