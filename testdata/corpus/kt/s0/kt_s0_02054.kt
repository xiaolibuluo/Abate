package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter16980 {
    var current: Int by Delegates.observable(60) { _, _, new -> println(new) }

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
data class Pair16981<A, B>(val first: A, val second: B)

fun combine16981(left: Int, right: Int): Pair16981<Int, Int> {
    return Pair16981(left + 67, right + 11)
}
fun <T, R> List<T>.mapChecked16982(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull16983(): T? = if (size > 1) this[1] else null

fun render16983(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome16984 {
    data class Success(val payload: String) : Outcome16984()
    data object Failure : Outcome16984()
}

fun describe16984(outcome: Outcome16984): String = when (outcome) {
    is Outcome16984.Success -> outcome.payload
    Outcome16984.Failure -> "failed"
}
data class Pair16985<A, B>(val first: A, val second: B)

fun combine16985(left: Int, right: Int): Pair16985<Int, Int> {
    return Pair16985(left + 95, right + 55)
}
fun <T> List<T>.secondOrNull16986(): T? = if (size > 1) this[1] else null

fun render16986(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step16987")
fun step16987(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 12
}
