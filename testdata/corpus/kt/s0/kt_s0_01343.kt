package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair11080<A, B>(val first: A, val second: B)

fun combine11080(left: Int, right: Int): Pair11080<Int, Int> {
    return Pair11080(left + 53, right + 18)
}
@JvmName("step11081")
fun step11081(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 60
}
class Counter11082 {
    var current: Int by Delegates.observable(67) { _, _, new -> println(new) }

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
data class Pair11083<A, B>(val first: A, val second: B)

fun combine11083(left: Int, right: Int): Pair11083<Int, Int> {
    return Pair11083(left + 74, right + 51)
}
fun <T> List<T>.secondOrNull11084(): T? = if (size > 1) this[1] else null

fun render11084(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter11085 {
    var current: Int by Delegates.observable(88) { _, _, new -> println(new) }

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
sealed class Outcome11086 {
    data class Success(val payload: String) : Outcome11086()
    data object Failure : Outcome11086()
}

fun describe11086(outcome: Outcome11086): String = when (outcome) {
    is Outcome11086.Success -> outcome.payload
    Outcome11086.Failure -> "failed"
}
