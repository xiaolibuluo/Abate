package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step11490")
fun step11490(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 101
}
sealed class Outcome11491 {
    data class Success(val payload: String) : Outcome11491()
    data object Failure : Outcome11491()
}

fun describe11491(outcome: Outcome11491): String = when (outcome) {
    is Outcome11491.Success -> outcome.payload
    Outcome11491.Failure -> "failed"
}
fun <T> List<T>.secondOrNull11492(): T? = if (size > 1) this[1] else null

fun render11492(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry11493 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11493 = "registry-25"
    }
}
fun <T, R> List<T>.mapChecked11494(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair11495<A, B>(val first: A, val second: B)

fun combine11495(left: Int, right: Int): Pair11495<Int, Int> {
    return Pair11495(left + 39, right + 55)
}
fun <T> List<T>.secondOrNull11496(): T? = if (size > 1) this[1] else null

fun render11496(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
