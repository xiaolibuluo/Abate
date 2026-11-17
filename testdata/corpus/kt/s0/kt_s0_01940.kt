package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked16020(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome16021 {
    data class Success(val payload: String) : Outcome16021()
    data object Failure : Outcome16021()
}

fun describe16021(outcome: Outcome16021): String = when (outcome) {
    is Outcome16021.Success -> outcome.payload
    Outcome16021.Failure -> "failed"
}
data class Pair16022<A, B>(val first: A, val second: B)

fun combine16022(left: Int, right: Int): Pair16022<Int, Int> {
    return Pair16022(left + 87, right + 21)
}
object Registry16023 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16023 = "registry-94"
    }
}
fun <T> List<T>.secondOrNull16024(): T? = if (size > 1) this[1] else null

fun render16024(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull16025(): T? = if (size > 1) this[1] else null

fun render16025(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry16026 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16026 = "registry-18"
    }
}
sealed class Outcome16027 {
    data class Success(val payload: String) : Outcome16027()
    data object Failure : Outcome16027()
}

fun describe16027(outcome: Outcome16027): String = when (outcome) {
    is Outcome16027.Success -> outcome.payload
    Outcome16027.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked16028(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
