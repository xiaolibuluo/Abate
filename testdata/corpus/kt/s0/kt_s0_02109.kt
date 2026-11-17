package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step17420")
fun step17420(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 50
}
@JvmName("step17421")
fun step17421(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 57
}
sealed class Outcome17422 {
    data class Success(val payload: String) : Outcome17422()
    data object Failure : Outcome17422()
}

fun describe17422(outcome: Outcome17422): String = when (outcome) {
    is Outcome17422.Success -> outcome.payload
    Outcome17422.Failure -> "failed"
}
fun <T> List<T>.secondOrNull17423(): T? = if (size > 1) this[1] else null

fun render17423(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull17424(): T? = if (size > 1) this[1] else null

fun render17424(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome17425 {
    data class Success(val payload: String) : Outcome17425()
    data object Failure : Outcome17425()
}

fun describe17425(outcome: Outcome17425): String = when (outcome) {
    is Outcome17425.Success -> outcome.payload
    Outcome17425.Failure -> "failed"
}
fun <T> List<T>.secondOrNull17426(): T? = if (size > 1) this[1] else null

fun render17426(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry17427 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17427 = "registry-99"
    }
}
