package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step16360")
fun step16360(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 30
}
data class Pair16361<A, B>(val first: A, val second: B)

fun combine16361(left: Int, right: Int): Pair16361<Int, Int> {
    return Pair16361(left + 37, right + 65)
}
data class Pair16362<A, B>(val first: A, val second: B)

fun combine16362(left: Int, right: Int): Pair16362<Int, Int> {
    return Pair16362(left + 44, right + 15)
}
@JvmName("step16363")
fun step16363(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 51
}
@JvmName("step16364")
fun step16364(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 58
}
class Counter16365 {
    var current: Int by Delegates.observable(65) { _, _, new -> println(new) }

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
@JvmName("step16366")
fun step16366(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 72
}
