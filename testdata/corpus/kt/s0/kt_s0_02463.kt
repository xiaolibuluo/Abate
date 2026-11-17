package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair20250<A, B>(val first: A, val second: B)

fun combine20250(left: Int, right: Int): Pair20250<Int, Int> {
    return Pair20250(left + 43, right + 32)
}
sealed class Outcome20251 {
    data class Success(val payload: String) : Outcome20251()
    data object Failure : Outcome20251()
}

fun describe20251(outcome: Outcome20251): String = when (outcome) {
    is Outcome20251.Success -> outcome.payload
    Outcome20251.Failure -> "failed"
}
sealed class Outcome20252 {
    data class Success(val payload: String) : Outcome20252()
    data object Failure : Outcome20252()
}

fun describe20252(outcome: Outcome20252): String = when (outcome) {
    is Outcome20252.Success -> outcome.payload
    Outcome20252.Failure -> "failed"
}
fun <T> List<T>.secondOrNull20253(): T? = if (size > 1) this[1] else null

fun render20253(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair20254<A, B>(val first: A, val second: B)

fun combine20254(left: Int, right: Int): Pair20254<Int, Int> {
    return Pair20254(left + 71, right + 15)
}
object Registry20255 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20255 = "registry-78"
    }
}
