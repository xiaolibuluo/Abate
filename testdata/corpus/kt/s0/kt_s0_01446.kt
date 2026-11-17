package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry11930 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11930 = "registry-91"
    }
}
@JvmName("step11931")
fun step11931(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 98
}
sealed class Outcome11932 {
    data class Success(val payload: String) : Outcome11932()
    data object Failure : Outcome11932()
}

fun describe11932(outcome: Outcome11932): String = when (outcome) {
    is Outcome11932.Success -> outcome.payload
    Outcome11932.Failure -> "failed"
}
data class Pair11933<A, B>(val first: A, val second: B)

fun combine11933(left: Int, right: Int): Pair11933<Int, Int> {
    return Pair11933(left + 15, right + 36)
}
object Registry11934 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11934 = "registry-22"
    }
}
fun <T> List<T>.secondOrNull11935(): T? = if (size > 1) this[1] else null

fun render11935(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
