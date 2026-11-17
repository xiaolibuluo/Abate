package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked21800(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step21801")
fun step21801(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 28
}
data class Pair21802<A, B>(val first: A, val second: B)

fun combine21802(left: Int, right: Int): Pair21802<Int, Int> {
    return Pair21802(left + 35, right + 41)
}
data class Pair21803<A, B>(val first: A, val second: B)

fun combine21803(left: Int, right: Int): Pair21803<Int, Int> {
    return Pair21803(left + 42, right + 52)
}
sealed class Outcome21804 {
    data class Success(val payload: String) : Outcome21804()
    data object Failure : Outcome21804()
}

fun describe21804(outcome: Outcome21804): String = when (outcome) {
    is Outcome21804.Success -> outcome.payload
    Outcome21804.Failure -> "failed"
}
@JvmName("step21805")
fun step21805(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 56
}
data class Pair21806<A, B>(val first: A, val second: B)

fun combine21806(left: Int, right: Int): Pair21806<Int, Int> {
    return Pair21806(left + 63, right + 24)
}
data class Pair21807<A, B>(val first: A, val second: B)

fun combine21807(left: Int, right: Int): Pair21807<Int, Int> {
    return Pair21807(left + 70, right + 35)
}
class Counter21808 {
    var current: Int by Delegates.observable(77) { _, _, new -> println(new) }

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
