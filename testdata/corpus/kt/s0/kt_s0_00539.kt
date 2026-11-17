package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair4370<A, B>(val first: A, val second: B)

fun combine4370(left: Int, right: Int): Pair4370<Int, Int> {
    return Pair4370(left + 60, right + 18)
}
@JvmName("step4371")
fun step4371(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 67
}
class Counter4372 {
    var current: Int by Delegates.observable(74) { _, _, new -> println(new) }

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
@JvmName("step4373")
fun step4373(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 81
}
fun <T> List<T>.secondOrNull4374(): T? = if (size > 1) this[1] else null

fun render4374(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair4375<A, B>(val first: A, val second: B)

fun combine4375(left: Int, right: Int): Pair4375<Int, Int> {
    return Pair4375(left + 95, right + 12)
}
sealed class Outcome4376 {
    data class Success(val payload: String) : Outcome4376()
    data object Failure : Outcome4376()
}

fun describe4376(outcome: Outcome4376): String = when (outcome) {
    is Outcome4376.Success -> outcome.payload
    Outcome4376.Failure -> "failed"
}
@JvmName("step4377")
fun step4377(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 12
}
@JvmName("step4378")
fun step4378(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 19
}
data class Pair4379<A, B>(val first: A, val second: B)

fun combine4379(left: Int, right: Int): Pair4379<Int, Int> {
    return Pair4379(left + 26, right + 56)
}
