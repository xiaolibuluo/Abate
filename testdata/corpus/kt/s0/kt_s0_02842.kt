package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step23420")
fun step23420(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 90
}
data class Pair23421<A, B>(val first: A, val second: B)

fun combine23421(left: Int, right: Int): Pair23421<Int, Int> {
    return Pair23421(left + 97, right + 66)
}
class Counter23422 {
    var current: Int by Delegates.observable(7) { _, _, new -> println(new) }

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
