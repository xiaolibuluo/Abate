package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair14450<A, B>(val first: A, val second: B)

fun combine14450(left: Int, right: Int): Pair14450<Int, Int> {
    return Pair14450(left + 69, right + 59)
}
object Registry14451 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14451 = "registry-76"
    }
}
fun <T> List<T>.secondOrNull14452(): T? = if (size > 1) this[1] else null

fun render14452(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry14453 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14453 = "registry-90"
    }
}
fun <T> List<T>.secondOrNull14454(): T? = if (size > 1) this[1] else null

fun render14454(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull14455(): T? = if (size > 1) this[1] else null

fun render14455(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome14456 {
    data class Success(val payload: String) : Outcome14456()
    data object Failure : Outcome14456()
}

fun describe14456(outcome: Outcome14456): String = when (outcome) {
    is Outcome14456.Success -> outcome.payload
    Outcome14456.Failure -> "failed"
}
