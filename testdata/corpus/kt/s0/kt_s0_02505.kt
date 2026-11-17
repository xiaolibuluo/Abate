package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair20610<A, B>(val first: A, val second: B)

fun combine20610(left: Int, right: Int): Pair20610<Int, Int> {
    return Pair20610(left + 26, right + 40)
}
@JvmName("step20611")
fun step20611(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 33
}
data class Pair20612<A, B>(val first: A, val second: B)

fun combine20612(left: Int, right: Int): Pair20612<Int, Int> {
    return Pair20612(left + 40, right + 62)
}
@JvmName("step20613")
fun step20613(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 47
}
class Counter20614 {
    var current: Int by Delegates.observable(54) { _, _, new -> println(new) }

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
