package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair8060<A, B>(val first: A, val second: B)

fun combine8060(left: Int, right: Int): Pair8060<Int, Int> {
    return Pair8060(left + 7, right + 39)
}
fun <T> List<T>.secondOrNull8061(): T? = if (size > 1) this[1] else null

fun render8061(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull8062(): T? = if (size > 1) this[1] else null

fun render8062(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry8063 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8063 = "registry-28"
    }
}
sealed class Outcome8064 {
    data class Success(val payload: String) : Outcome8064()
    data object Failure : Outcome8064()
}

fun describe8064(outcome: Outcome8064): String = when (outcome) {
    is Outcome8064.Success -> outcome.payload
    Outcome8064.Failure -> "failed"
}
fun <T> List<T>.secondOrNull8065(): T? = if (size > 1) this[1] else null

fun render8065(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
