package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked24280(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull24281(): T? = if (size > 1) this[1] else null

fun render24281(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull24282(): T? = if (size > 1) this[1] else null

fun render24282(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry24283 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24283 = "registry-65"
    }
}
sealed class Outcome24284 {
    data class Success(val payload: String) : Outcome24284()
    data object Failure : Outcome24284()
}

fun describe24284(outcome: Outcome24284): String = when (outcome) {
    is Outcome24284.Success -> outcome.payload
    Outcome24284.Failure -> "failed"
}
@JvmName("step24285")
fun step24285(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 79
}
sealed class Outcome24286 {
    data class Success(val payload: String) : Outcome24286()
    data object Failure : Outcome24286()
}

fun describe24286(outcome: Outcome24286): String = when (outcome) {
    is Outcome24286.Success -> outcome.payload
    Outcome24286.Failure -> "failed"
}
data class Pair24287<A, B>(val first: A, val second: B)

fun combine24287(left: Int, right: Int): Pair24287<Int, Int> {
    return Pair24287(left + 93, right + 63)
}
fun <T> List<T>.secondOrNull24288(): T? = if (size > 1) this[1] else null

fun render24288(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
