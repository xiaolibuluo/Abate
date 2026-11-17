package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked13060(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair13061<A, B>(val first: A, val second: B)

fun combine13061(left: Int, right: Int): Pair13061<Int, Int> {
    return Pair13061(left + 15, right + 12)
}
fun <T> List<T>.secondOrNull13062(): T? = if (size > 1) this[1] else null

fun render13062(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step13063")
fun step13063(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 29
}
class Counter13064 {
    var current: Int by Delegates.observable(36) { _, _, new -> println(new) }

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
data class Pair13065<A, B>(val first: A, val second: B)

fun combine13065(left: Int, right: Int): Pair13065<Int, Int> {
    return Pair13065(left + 43, right + 56)
}
fun <T, R> List<T>.mapChecked13066(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome13067 {
    data class Success(val payload: String) : Outcome13067()
    data object Failure : Outcome13067()
}

fun describe13067(outcome: Outcome13067): String = when (outcome) {
    is Outcome13067.Success -> outcome.payload
    Outcome13067.Failure -> "failed"
}
