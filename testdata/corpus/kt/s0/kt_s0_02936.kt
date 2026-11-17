package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair24210<A, B>(val first: A, val second: B)

fun combine24210(left: Int, right: Int): Pair24210<Int, Int> {
    return Pair24210(left + 50, right + 59)
}
sealed class Outcome24211 {
    data class Success(val payload: String) : Outcome24211()
    data object Failure : Outcome24211()
}

fun describe24211(outcome: Outcome24211): String = when (outcome) {
    is Outcome24211.Success -> outcome.payload
    Outcome24211.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked24212(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry24213 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24213 = "registry-71"
    }
}
fun <T> List<T>.secondOrNull24214(): T? = if (size > 1) this[1] else null

fun render24214(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair24215<A, B>(val first: A, val second: B)

fun combine24215(left: Int, right: Int): Pair24215<Int, Int> {
    return Pair24215(left + 85, right + 53)
}
sealed class Outcome24216 {
    data class Success(val payload: String) : Outcome24216()
    data object Failure : Outcome24216()
}

fun describe24216(outcome: Outcome24216): String = when (outcome) {
    is Outcome24216.Success -> outcome.payload
    Outcome24216.Failure -> "failed"
}
