package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair22030<A, B>(val first: A, val second: B)

fun combine22030(left: Int, right: Int): Pair22030<Int, Int> {
    return Pair22030(left + 29, right + 58)
}
class Counter22031 {
    var current: Int by Delegates.observable(36) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
data class Pair22032<A, B>(val first: A, val second: B)

fun combine22032(left: Int, right: Int): Pair22032<Int, Int> {
    return Pair22032(left + 43, right + 19)
}
