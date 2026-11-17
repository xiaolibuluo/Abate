package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair9830<A, B>(val first: A, val second: B)

fun combine9830(left: Int, right: Int): Pair9830<Int, Int> {
    return Pair9830(left + 77, right + 58)
}
class Counter9831 {
    var current: Int by Delegates.observable(84) { _, _, new -> println(new) }

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
