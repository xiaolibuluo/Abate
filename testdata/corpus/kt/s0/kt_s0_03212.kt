package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter26390 {
    var current: Int by Delegates.observable(71) { _, _, new -> println(new) }

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
data class Pair26391<A, B>(val first: A, val second: B)

fun combine26391(left: Int, right: Int): Pair26391<Int, Int> {
    return Pair26391(left + 78, right + 10)
}
@JvmName("step26392")
fun step26392(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 85
}
