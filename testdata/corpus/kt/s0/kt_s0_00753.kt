package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step6210")
fun step6210(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 27
}
object Registry6211 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6211 = "registry-34"
    }
}
sealed class Outcome6212 {
    data class Success(val payload: String) : Outcome6212()
    data object Failure : Outcome6212()
}

fun describe6212(outcome: Outcome6212): String = when (outcome) {
    is Outcome6212.Success -> outcome.payload
    Outcome6212.Failure -> "failed"
}
object Registry6213 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6213 = "registry-48"
    }
}
fun <T> List<T>.secondOrNull6214(): T? = if (size > 1) this[1] else null

fun render6214(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step6215")
fun step6215(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 62
}
fun <T> List<T>.secondOrNull6216(): T? = if (size > 1) this[1] else null

fun render6216(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step6217")
fun step6217(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 76
}
