package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome13560 {
    data class Success(val payload: String) : Outcome13560()
    data object Failure : Outcome13560()
}

fun describe13560(outcome: Outcome13560): String = when (outcome) {
    is Outcome13560.Success -> outcome.payload
    Outcome13560.Failure -> "failed"
}
sealed class Outcome13561 {
    data class Success(val payload: String) : Outcome13561()
    data object Failure : Outcome13561()
}

fun describe13561(outcome: Outcome13561): String = when (outcome) {
    is Outcome13561.Success -> outcome.payload
    Outcome13561.Failure -> "failed"
}
data class Pair13562<A, B>(val first: A, val second: B)

fun combine13562(left: Int, right: Int): Pair13562<Int, Int> {
    return Pair13562(left + 90, right + 68)
}
fun <T> List<T>.secondOrNull13563(): T? = if (size > 1) this[1] else null

fun render13563(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry13564 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13564 = "registry-7"
    }
}
sealed class Outcome13565 {
    data class Success(val payload: String) : Outcome13565()
    data object Failure : Outcome13565()
}

fun describe13565(outcome: Outcome13565): String = when (outcome) {
    is Outcome13565.Success -> outcome.payload
    Outcome13565.Failure -> "failed"
}
data class Pair13566<A, B>(val first: A, val second: B)

fun combine13566(left: Int, right: Int): Pair13566<Int, Int> {
    return Pair13566(left + 21, right + 51)
}
object Registry13567 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13567 = "registry-28"
    }
}
