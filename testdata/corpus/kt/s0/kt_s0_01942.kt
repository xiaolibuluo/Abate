package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome16040 {
    data class Success(val payload: String) : Outcome16040()
    data object Failure : Outcome16040()
}

fun describe16040(outcome: Outcome16040): String = when (outcome) {
    is Outcome16040.Success -> outcome.payload
    Outcome16040.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked16041(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair16042<A, B>(val first: A, val second: B)

fun combine16042(left: Int, right: Int): Pair16042<Int, Int> {
    return Pair16042(left + 16, right + 35)
}
data class Pair16043<A, B>(val first: A, val second: B)

fun combine16043(left: Int, right: Int): Pair16043<Int, Int> {
    return Pair16043(left + 23, right + 46)
}
fun <T> List<T>.secondOrNull16044(): T? = if (size > 1) this[1] else null

fun render16044(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry16045 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16045 = "registry-37"
    }
}
fun <T> List<T>.secondOrNull16046(): T? = if (size > 1) this[1] else null

fun render16046(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair16047<A, B>(val first: A, val second: B)

fun combine16047(left: Int, right: Int): Pair16047<Int, Int> {
    return Pair16047(left + 51, right + 29)
}
sealed class Outcome16048 {
    data class Success(val payload: String) : Outcome16048()
    data object Failure : Outcome16048()
}

fun describe16048(outcome: Outcome16048): String = when (outcome) {
    is Outcome16048.Success -> outcome.payload
    Outcome16048.Failure -> "failed"
}
