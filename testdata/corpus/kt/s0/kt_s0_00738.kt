package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry6080 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6080 = "registry-52"
    }
}
@JvmName("step6081")
fun step6081(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 59
}
object Registry6082 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6082 = "registry-66"
    }
}
fun <T> List<T>.secondOrNull6083(): T? = if (size > 1) this[1] else null

fun render6083(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry6084 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6084 = "registry-80"
    }
}
sealed class Outcome6085 {
    data class Success(val payload: String) : Outcome6085()
    data object Failure : Outcome6085()
}

fun describe6085(outcome: Outcome6085): String = when (outcome) {
    is Outcome6085.Success -> outcome.payload
    Outcome6085.Failure -> "failed"
}
