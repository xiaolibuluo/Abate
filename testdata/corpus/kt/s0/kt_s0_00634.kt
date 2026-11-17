package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome5180 {
    data class Success(val payload: String) : Outcome5180()
    data object Failure : Outcome5180()
}

fun describe5180(outcome: Outcome5180): String = when (outcome) {
    is Outcome5180.Success -> outcome.payload
    Outcome5180.Failure -> "failed"
}
data class Pair5181<A, B>(val first: A, val second: B)

fun combine5181(left: Int, right: Int): Pair5181<Int, Int> {
    return Pair5181(left + 53, right + 47)
}
fun <T> List<T>.secondOrNull5182(): T? = if (size > 1) this[1] else null

fun render5182(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull5183(): T? = if (size > 1) this[1] else null

fun render5183(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry5184 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5184 = "registry-74"
    }
}
fun <T> List<T>.secondOrNull5185(): T? = if (size > 1) this[1] else null

fun render5185(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair5186<A, B>(val first: A, val second: B)

fun combine5186(left: Int, right: Int): Pair5186<Int, Int> {
    return Pair5186(left + 88, right + 41)
}
