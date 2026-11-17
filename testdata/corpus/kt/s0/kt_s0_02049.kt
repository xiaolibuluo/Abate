package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair16940<A, B>(val first: A, val second: B)

fun combine16940(left: Int, right: Int): Pair16940<Int, Int> {
    return Pair16940(left + 8, right + 33)
}
fun <T> List<T>.secondOrNull16941(): T? = if (size > 1) this[1] else null

fun render16941(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair16942<A, B>(val first: A, val second: B)

fun combine16942(left: Int, right: Int): Pair16942<Int, Int> {
    return Pair16942(left + 22, right + 55)
}
fun <T> List<T>.secondOrNull16943(): T? = if (size > 1) this[1] else null

fun render16943(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull16944(): T? = if (size > 1) this[1] else null

fun render16944(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull16945(): T? = if (size > 1) this[1] else null

fun render16945(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step16946")
fun step16946(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 50
}
fun <T> List<T>.secondOrNull16947(): T? = if (size > 1) this[1] else null

fun render16947(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair16948<A, B>(val first: A, val second: B)

fun combine16948(left: Int, right: Int): Pair16948<Int, Int> {
    return Pair16948(left + 64, right + 60)
}
