package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step10480")
fun step10480(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 49
}
sealed class Outcome10481 {
    data class Success(val payload: String) : Outcome10481()
    data object Failure : Outcome10481()
}

fun describe10481(outcome: Outcome10481): String = when (outcome) {
    is Outcome10481.Success -> outcome.payload
    Outcome10481.Failure -> "failed"
}
fun <T> List<T>.secondOrNull10482(): T? = if (size > 1) this[1] else null

fun render10482(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry10483 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10483 = "registry-70"
    }
}
fun <T, R> List<T>.mapChecked10484(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step10485")
fun step10485(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 84
}
