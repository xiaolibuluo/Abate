package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step14890")
fun step14890(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 59
}
fun <T> List<T>.secondOrNull14891(): T? = if (size > 1) this[1] else null

fun render14891(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome14892 {
    data class Success(val payload: String) : Outcome14892()
    data object Failure : Outcome14892()
}

fun describe14892(outcome: Outcome14892): String = when (outcome) {
    is Outcome14892.Success -> outcome.payload
    Outcome14892.Failure -> "failed"
}
fun <T> List<T>.secondOrNull14893(): T? = if (size > 1) this[1] else null

fun render14893(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked14894(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry14895 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14895 = "registry-94"
    }
}
