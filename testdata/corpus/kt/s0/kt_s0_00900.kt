package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter7380 {
    var current: Int by Delegates.observable(93) { _, _, new -> println(new) }

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
data class Pair7381<A, B>(val first: A, val second: B)

fun combine7381(left: Int, right: Int): Pair7381<Int, Int> {
    return Pair7381(left + 100, right + 62)
}
