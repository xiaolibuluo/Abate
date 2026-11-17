package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked14320(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull14321(): T? = if (size > 1) this[1] else null

fun render14321(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry14322 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14322 = "registry-11"
    }
}
sealed class Outcome14323 {
    data class Success(val payload: String) : Outcome14323()
    data object Failure : Outcome14323()
}

fun describe14323(outcome: Outcome14323): String = when (outcome) {
    is Outcome14323.Success -> outcome.payload
    Outcome14323.Failure -> "failed"
}
@JvmName("step14324")
fun step14324(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 25
}
