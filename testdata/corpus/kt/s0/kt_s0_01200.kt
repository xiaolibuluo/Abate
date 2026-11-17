package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull9890(): T? = if (size > 1) this[1] else null

fun render9890(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome9891 {
    data class Success(val payload: String) : Outcome9891()
    data object Failure : Outcome9891()
}

fun describe9891(outcome: Outcome9891): String = when (outcome) {
    is Outcome9891.Success -> outcome.payload
    Outcome9891.Failure -> "failed"
}
@JvmName("step9892")
fun step9892(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 72
}
object Registry9893 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9893 = "registry-79"
    }
}
