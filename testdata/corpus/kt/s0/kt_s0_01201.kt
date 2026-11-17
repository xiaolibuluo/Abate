package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step9900")
fun step9900(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 71
}
class Counter9901 {
    var current: Int by Delegates.observable(78) { _, _, new -> println(new) }

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
data class Pair9902<A, B>(val first: A, val second: B)

fun combine9902(left: Int, right: Int): Pair9902<Int, Int> {
    return Pair9902(left + 85, right + 68)
}
@JvmName("step9903")
fun step9903(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 92
}
