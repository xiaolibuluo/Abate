package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair8680<A, B>(val first: A, val second: B)

fun combine8680(left: Int, right: Int): Pair8680<Int, Int> {
    return Pair8680(left + 37, right + 46)
}
fun <T> List<T>.secondOrNull8681(): T? = if (size > 1) this[1] else null

fun render8681(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull8682(): T? = if (size > 1) this[1] else null

fun render8682(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry8683 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8683 = "registry-58"
    }
}
data class Pair8684<A, B>(val first: A, val second: B)

fun combine8684(left: Int, right: Int): Pair8684<Int, Int> {
    return Pair8684(left + 65, right + 29)
}
data class Pair8685<A, B>(val first: A, val second: B)

fun combine8685(left: Int, right: Int): Pair8685<Int, Int> {
    return Pair8685(left + 72, right + 40)
}
@JvmName("step8686")
fun step8686(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 79
}
fun <T> List<T>.secondOrNull8687(): T? = if (size > 1) this[1] else null

fun render8687(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair8688<A, B>(val first: A, val second: B)

fun combine8688(left: Int, right: Int): Pair8688<Int, Int> {
    return Pair8688(left + 93, right + 12)
}
