package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step14980")
fun step14980(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 79
}
object Registry14981 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14981 = "registry-86"
    }
}
sealed class Outcome14982 {
    data class Success(val payload: String) : Outcome14982()
    data object Failure : Outcome14982()
}

fun describe14982(outcome: Outcome14982): String = when (outcome) {
    is Outcome14982.Success -> outcome.payload
    Outcome14982.Failure -> "failed"
}
object Registry14983 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14983 = "registry-100"
    }
}
fun <T> List<T>.secondOrNull14984(): T? = if (size > 1) this[1] else null

fun render14984(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull14985(): T? = if (size > 1) this[1] else null

fun render14985(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome14986 {
    data class Success(val payload: String) : Outcome14986()
    data object Failure : Outcome14986()
}

fun describe14986(outcome: Outcome14986): String = when (outcome) {
    is Outcome14986.Success -> outcome.payload
    Outcome14986.Failure -> "failed"
}
