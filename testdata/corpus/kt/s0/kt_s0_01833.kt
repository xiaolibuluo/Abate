package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair15060<A, B>(val first: A, val second: B)

fun combine15060(left: Int, right: Int): Pair15060<Int, Int> {
    return Pair15060(left + 86, right + 59)
}
fun <T, R> List<T>.mapChecked15061(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull15062(): T? = if (size > 1) this[1] else null

fun render15062(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull15063(): T? = if (size > 1) this[1] else null

fun render15063(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome15064 {
    data class Success(val payload: String) : Outcome15064()
    data object Failure : Outcome15064()
}

fun describe15064(outcome: Outcome15064): String = when (outcome) {
    is Outcome15064.Success -> outcome.payload
    Outcome15064.Failure -> "failed"
}
object Registry15065 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15065 = "registry-24"
    }
}
data class Pair15066<A, B>(val first: A, val second: B)

fun combine15066(left: Int, right: Int): Pair15066<Int, Int> {
    return Pair15066(left + 31, right + 64)
}
fun <T> List<T>.secondOrNull15067(): T? = if (size > 1) this[1] else null

fun render15067(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull15068(): T? = if (size > 1) this[1] else null

fun render15068(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
