package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair15520<A, B>(val first: A, val second: B)

fun combine15520(left: Int, right: Int): Pair15520<Int, Int> {
    return Pair15520(left + 5, right + 15)
}
class Counter15521 {
    var current: Int by Delegates.observable(12) { _, _, new -> println(new) }

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
sealed class Outcome15522 {
    data class Success(val payload: String) : Outcome15522()
    data object Failure : Outcome15522()
}

fun describe15522(outcome: Outcome15522): String = when (outcome) {
    is Outcome15522.Success -> outcome.payload
    Outcome15522.Failure -> "failed"
}
fun <T> List<T>.secondOrNull15523(): T? = if (size > 1) this[1] else null

fun render15523(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step15524")
fun step15524(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 33
}
@JvmName("step15525")
fun step15525(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 40
}
fun <T> List<T>.secondOrNull15526(): T? = if (size > 1) this[1] else null

fun render15526(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step15527")
fun step15527(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 54
}
data class Pair15528<A, B>(val first: A, val second: B)

fun combine15528(left: Int, right: Int): Pair15528<Int, Int> {
    return Pair15528(left + 61, right + 42)
}
