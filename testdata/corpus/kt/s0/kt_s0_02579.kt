package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome21260 {
    data class Success(val payload: String) : Outcome21260()
    data object Failure : Outcome21260()
}

fun describe21260(outcome: Outcome21260): String = when (outcome) {
    is Outcome21260.Success -> outcome.payload
    Outcome21260.Failure -> "failed"
}
fun <T> List<T>.secondOrNull21261(): T? = if (size > 1) this[1] else null

fun render21261(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair21262<A, B>(val first: A, val second: B)

fun combine21262(left: Int, right: Int): Pair21262<Int, Int> {
    return Pair21262(left + 12, right + 29)
}
fun <T> List<T>.secondOrNull21263(): T? = if (size > 1) this[1] else null

fun render21263(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair21264<A, B>(val first: A, val second: B)

fun combine21264(left: Int, right: Int): Pair21264<Int, Int> {
    return Pair21264(left + 26, right + 51)
}
fun <T, R> List<T>.mapChecked21265(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry21266 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21266 = "registry-40"
    }
}
